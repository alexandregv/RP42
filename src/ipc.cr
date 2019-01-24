require "socket"

module RichCrystal
  class Ipc
    # Enumerate the differents opcodes
    enum Opcode : Int32
      Handshake = 0
      Frame = 1
    end

    # Create the UNIXSocket with ipc_path and socket name 'discord-ipc-0'
    def initialize
      @socket = UNIXSocket.new("#{ipc_path}discord-ipc-0")
    end

    # Return where is the discord-ipc-0 socket with
    # different environment variables
    def ipc_path : String
      # Possibles path environment variables names
      variables = %w(XDG_RUNTIME_DIR TMPDIR TMP TEMP)

      # Iterate environment variables names
      variables.each do |variable_name|
      # Handling a key error if the environment variable does not exists
        begin
          variable = ENV[variable_name]
          return variable
        rescue KeyError
          next # Continue the loop if the variable does not exists
        end
      end

      # If none of the environment variables have been found return '\tmp'
      "\tmp"
    end

    # Send a payload to the UNIXSocket with the opcode, returns
    # the JSON response
    def send(opcode : Opcode, payload : String) : String
      # Write the opcode and the payload size as LitteEndian
      @socket.write_bytes(opcode.value, IO::ByteFormat::LittleEndian)
      @socket.write_bytes(payload.size, IO::ByteFormat::LittleEndian)

      # And then add the payload
      @socket << payload

      # Read the response code
      code = @socket.read_bytes(Int32, IO::ByteFormat::LittleEndian)
      # Read the size of the data
      data_size = @socket.read_bytes(Int32, IO::ByteFormat::LittleEndian)

      # Then fully read the data_size number of bytes and convert them into a String
      bytes = Bytes.new(data_size)
      @socket.read_fully(bytes)
      String.new(bytes)
    end
  end
end
