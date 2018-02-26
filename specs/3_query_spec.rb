require 'rspec'
require 'httpclient'
require 'json'
require 'google/protobuf'
require 'securerandom'

shared_examples "get tariff" do |hoursParked, expectedPrice|
  context "Ticket issued just under #{hoursParked} hours ago" do
    before :all do
      # create a ticket less than an hour old
      @ticketID = SecureRandom.uuid
      timestamp = Google::Protobuf::Timestamp.new(:seconds => Time.now.to_i - (hoursParked * 60 * 59))
      ticket = Events::TicketIssued.new(:at => timestamp, :ticketID => @ticketID)
      encoded_data = Events::TicketIssued.encode(ticket)
      $pg.exec_params("INSERT INTO events (type, data) VALUES ($1,$2)", ["TicketIssued", $pg.escape_bytea(encoded_data)])

      @response = $client.get "#{$endpoint}/ticket/#{@ticketID}/tariff"
    end

    it "returns 200 OK" do
      expect(@response.status).to eq 200
    end

    it "returns the ticket ID" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["id"]).to eq @ticketID
    end

    it "returns a price of £#{sprintf('%.2f', expectedPrice/100.0)}" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["tariff"]).to eq expectedPrice
    end
  end
end


describe "GET /ticket/{id}/tariff :: Get the current price of the ticket" do
  after :all do
    $pg.exec("TRUNCATE events")
  end

  context "Unknown ticket" do
    before :all do
      lessThanHourTicketID = 'unknown-ticket'
      @response = $client.get "#{$endpoint}/ticket/#{lessThanHourTicketID}/tariff"
    end

    it "returns 404 Not Found" do
      expect(@response.status).to eq 404
    end
  end

  include_examples "get tariff", 1, 150
  include_examples "get tariff", 3, 300
  include_examples "get tariff", 6, 1000
  include_examples "get tariff", 24, 2000
  include_examples "get tariff", 48, 5000
  include_examples "get tariff", 72, 7500
  include_examples "get tariff", 96, 10000
end
