require 'redis'
require 'net/http'
require 'json'
require 'byebug'

r = Redis.new

facts = [
  {"id" => 1, "type" => "A", "sku" => "E031", "amount" => 100},
  {"id" => 2, "type" => "B", "sku" => "E032", "amount" => 200},
  {"id" => 3, "type" => "C", "sku" => "E033", "amount" => 300},
  {"id" => 4, "type" => "D", "sku" => "E034", "amount" => 400},
]

r.set 1, facts.to_json

uri = URI.parse("http://localhost:9999")
response = Net::HTTP.post_form(uri, {"id" => 1})
puts response.body
