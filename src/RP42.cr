require "./client"

module RP42
	VERSION = "0.1.0"

  hostname = System.hostname.split(".", 2)
  exit if hostname[1] != "42.fr" 

  client = RichCrystal::Client.new(531103976029028367_u64)
	client.login
	client.activity({
		"details" => "Location: #{hostname[0]}",
    "assets" => {
      "large_image" => "logo",
      "large_text" => "42",
    }
	})

  sleep
end
