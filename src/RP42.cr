require "oauth2"
require "./client"

module RP42
	VERSION = "0.1.0"

  api_id = "6006639307c5ecc1506feb96d9f833df325e07573040acecad429f75b824a72e"
  api_secret = "<API-SECRET>"

  oauth2_client = OAuth2::Client.new("api.intra.42.fr", api_id, api_secret, authorize_uri: "/oauth/authorize", token_uri: "/oauth/token")
  access_token = oauth2_client.get_access_token_using_client_credentials
  http_client = HTTP::Client.new("api.intra.42.fr", tls: OpenSSL::SSL::Context::Client.insecure)
  access_token.authenticate(http_client)

  username = `whoami`.chomp("\n")
  hostname = System.hostname.split(".", 2)
  hostname[0] = "At home" if hostname[1] != "42.fr"

  coa = JSON.parse(http_client.get("/v2/users/#{username}/coalitions").body)[1]["name"].to_s
  lvl = JSON.parse(http_client.get("/v2/users/#{username}").body)["cursus_users"][0]["level"].to_s
  
  rich_client = RichCrystal::Client.new(531103976029028367_u64)
    rich_client.login
    rich_client.activity({
    	"details" => "Level: #{lvl}",
    	"state" => "Location: #{hostname[0]}",
    "assets" => {
      "large_image" => "logo",
      "large_text" => username,
      "small_image" => coa.tr(" ", "-").downcase,
      "small_text" => coa,
    }
  })

  sleep
end
