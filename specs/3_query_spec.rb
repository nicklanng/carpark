require 'rspec'
require 'httpclient'
require 'json'
require 'google/protobuf'
require 'securerandom'

# TODO: Add helper method to these tests that takes a time and a price

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

  context "Known ticket under an hour" do
    before :all do
      # create a ticket less than an hour old
      @lessThanHourTicketID = SecureRandom.uuid
      timestamp = Google::Protobuf::Timestamp.new(:seconds => Time.now.to_i)
      knownTicket = Events::TicketIssued.new(:at => timestamp, :ticketID => @lessThanHourTicketID)
      encoded_data = Events::TicketIssued.encode(knownTicket)
      $pg.exec_params("INSERT INTO events (type, data) VALUES ($1,$2)", ["TicketIssued", $pg.escape_bytea(encoded_data)])

      @response = $client.get "#{$endpoint}/ticket/#{@lessThanHourTicketID}/tariff"
    end

    it "returns 200 OK" do
      expect(@response.status).to eq 200
    end

    it "returns the ticket ID" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["id"]).to eq @lessThanHourTicketID
    end

    it "returns a price of £1.50" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["tariff"]).to eq 150
    end
  end

  context "Known ticket under 3 hours" do
    before :all do
      # create a ticket less than 3 hours old
      @lessThan3HourTicketID = SecureRandom.uuid
      timestamp = Google::Protobuf::Timestamp.new(:seconds => Time.now.to_i - (60 * 60 * 2))
      knownTicket = Events::TicketIssued.new(:at => timestamp, :ticketID => @lessThan3HourTicketID)
      encoded_data = Events::TicketIssued.encode(knownTicket)
      $pg.exec_params("INSERT INTO events (type, data) VALUES ($1,$2)", ["TicketIssued", $pg.escape_bytea(encoded_data)])

      @response = $client.get "#{$endpoint}/ticket/#{@lessThan3HourTicketID}/tariff"
    end

    it "returns 200 OK" do
      expect(@response.status).to eq 200
    end

    it "returns the ticket ID" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["id"]).to eq @lessThan3HourTicketID
    end

    it "returns a price of £3.00" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["tariff"]).to eq 300
    end
  end

  context "Known ticket under 6 hours" do
    before :all do
      # create a ticket less than 6 hours old
      @lessThan6HourTicketID = SecureRandom.uuid
      timestamp = Google::Protobuf::Timestamp.new(:seconds => Time.now.to_i - (60 * 60 * 5))
      knownTicket = Events::TicketIssued.new(:at => timestamp, :ticketID => @lessThan6HourTicketID)
      encoded_data = Events::TicketIssued.encode(knownTicket)
      $pg.exec_params("INSERT INTO events (type, data) VALUES ($1,$2)", ["TicketIssued", $pg.escape_bytea(encoded_data)])

      @response = $client.get "#{$endpoint}/ticket/#{@lessThan6HourTicketID}/tariff"
    end

    it "returns 200 OK" do
      expect(@response.status).to eq 200
    end

    it "returns the ticket ID" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["id"]).to eq @lessThan6HourTicketID
    end

    it "returns a price of £10.00" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["tariff"]).to eq 1000
    end
  end

  context "Known ticket under 24 hours" do
    before :all do
      # create a ticket less than 24 hours old
      @lessThan24HourTicketID = SecureRandom.uuid
      timestamp = Google::Protobuf::Timestamp.new(:seconds => Time.now.to_i - (60 * 60 * 23))
      knownTicket = Events::TicketIssued.new(:at => timestamp, :ticketID => @lessThan24HourTicketID)
      encoded_data = Events::TicketIssued.encode(knownTicket)
      $pg.exec_params("INSERT INTO events (type, data) VALUES ($1,$2)", ["TicketIssued", $pg.escape_bytea(encoded_data)])

      @response = $client.get "#{$endpoint}/ticket/#{@lessThan24HourTicketID}/tariff"
    end

    it "returns 200 OK" do
      expect(@response.status).to eq 200
    end

    it "returns the ticket ID" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["id"]).to eq @lessThan24HourTicketID
    end

    it "returns a price of £20.00" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["tariff"]).to eq 2000
    end
  end

  context "Known ticket under 48 hours" do
    before :all do
      # create a ticket less than 48 hours old
      @lessThan48HourTicketID = SecureRandom.uuid
      timestamp = Google::Protobuf::Timestamp.new(:seconds => Time.now.to_i - (60 * 60 * 47))
      knownTicket = Events::TicketIssued.new(:at => timestamp, :ticketID => @lessThan48HourTicketID)
      encoded_data = Events::TicketIssued.encode(knownTicket)
      $pg.exec_params("INSERT INTO events (type, data) VALUES ($1,$2)", ["TicketIssued", $pg.escape_bytea(encoded_data)])

      @response = $client.get "#{$endpoint}/ticket/#{@lessThan48HourTicketID}/tariff"
    end

    it "returns 200 OK" do
      expect(@response.status).to eq 200
    end

    it "returns the ticket ID" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["id"]).to eq @lessThan48HourTicketID
    end

    it "returns a price of £50.00" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["tariff"]).to eq 5000
    end
  end

  context "Known ticket under 72 hours" do
    before :all do
      # create a ticket less than 72 hours old
      @lessThan72HourTicketID = SecureRandom.uuid
      timestamp = Google::Protobuf::Timestamp.new(:seconds => Time.now.to_i - (60 * 60 * 71))
      knownTicket = Events::TicketIssued.new(:at => timestamp, :ticketID => @lessThan72HourTicketID)
      encoded_data = Events::TicketIssued.encode(knownTicket)
      $pg.exec_params("INSERT INTO events (type, data) VALUES ($1,$2)", ["TicketIssued", $pg.escape_bytea(encoded_data)])

      @response = $client.get "#{$endpoint}/ticket/#{@lessThan72HourTicketID}/tariff"
    end

    it "returns 200 OK" do
      expect(@response.status).to eq 200
    end

    it "returns the ticket ID" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["id"]).to eq @lessThan72HourTicketID
    end

    it "returns a price of £75.00" do
      jsonResponse = JSON.parse @response.body
      expect(jsonResponse["tariff"]).to eq 7500
    end
  end
end
