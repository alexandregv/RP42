require "json"
require "./ipc"

module RichCrystal
  class Client
    # Creates a new Rich-crystal client used for set rich presence activity
    def initialize(@client_id : UInt64)
      @ipc = RichCrystal::Ipc.new
    end

    # Log the RichCrystal client by sending a first handshake to the IPC
    def login
      # Generate the payload in JSON
      payload = {
        "v"         => 1,
        "client_id" => "#{@client_id}",
        "nonce"     => Time.now.to_s("%s"),
      }.to_json

      # Send the handshake to the IPC
      @ipc.send(RichCrystal::Ipc::Opcode::Handshake, payload)
    end

    # Retrieves a Hash of Strings for sending the frame payload to the IPC
    # with discord-rich-presence parameters (see here https://github.com/discordapp/discord-rpc/blob/master/documentation/hard-mode.md#new-rpc-command)
    # and return the JSON response
    def activity(activity)
      # Generate the payload in JSON
      payload = {
        "cmd"  => "SET_ACTIVITY",
        "args" => {
          "pid"      => Process.pid,
          "activity" => activity,
        },
        "nonce" => Time.now.to_s("%s"),
      }.to_json

      # Send the frame to the IPC
      @ipc.send(RichCrystal::Ipc::Opcode::Frame, payload)
    end
  end
end
