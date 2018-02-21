require 'rspec'
require 'httpclient'
require 'json'

describe "GET /_status :: Health Check is not yet implemented" do
  before :all do
    @response = $client.get "#{$endpoint}/_status"
  end

  it "returns 501 NOT IMPLEMENTED" do
    expect(@response.status).to eq(501)
  end
end
