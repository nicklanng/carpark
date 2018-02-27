require 'rspec'
require 'httpclient'
require 'json'
require 'google/protobuf'

describe "POST /ticket/{ID}/complete :: Use a ticket to leave the car park" do
  context "known ticket that has only been issued" do
    before :all do
      @ticketID = SecureRandom.uuid
      timestamp = Google::Protobuf::Timestamp.new(:seconds => Time.now.to_i)
      ticket = Events::TicketIssued.new(:at => timestamp, :ticketID => @ticketID)
      encoded_data = Events::TicketIssued.encode(ticket)
      res = $pg.exec_params("INSERT INTO events (type, data) VALUES ($1,$2)", ["TicketIssued", $pg.escape_bytea(encoded_data)])

      @requestSent = Time.now.to_i
      @response = $client.post "#{$endpoint}/ticket/#{@ticketID}/complete"
    end

    it "returns 400 BAD REQUEST" do
      expect(@response.status).to eq 400
    end
  end

  context "known ticket that has already been paid" do
    before :all do
      @ticketID = SecureRandom.uuid
      timestamp = Google::Protobuf::Timestamp.new(:seconds => Time.now.to_i)
      ticket = Events::TicketIssued.new(:at => timestamp, :ticketID => @ticketID)
      encoded_data = Events::TicketIssued.encode(ticket)
      res = $pg.exec_params("INSERT INTO events (type, data) VALUES ($1,$2)", ["TicketIssued", $pg.escape_bytea(encoded_data)])

      ticket = Events::TicketPaid.new(:at => timestamp, :ticketID => @ticketID)
      encoded_data = Events::TicketPaid.encode(ticket)
      res = $pg.exec_params("INSERT INTO events (type, data) VALUES ($1,$2)", ["TicketPaid", $pg.escape_bytea(encoded_data)])

      @requestSent = Time.now.to_i
      @response = $client.post "#{$endpoint}/ticket/#{@ticketID}/complete"
    end

    it "returns 200 OK" do
      expect(@response.status).to eq 200
    end
  end

  context "known ticket that has already been completed" do
    before :all do
      @ticketID = SecureRandom.uuid
      timestamp = Google::Protobuf::Timestamp.new(:seconds => Time.now.to_i)
      ticket = Events::TicketIssued.new(:at => timestamp, :ticketID => @ticketID)
      encoded_data = Events::TicketIssued.encode(ticket)
      res = $pg.exec_params("INSERT INTO events (type, data) VALUES ($1,$2)", ["TicketIssued", $pg.escape_bytea(encoded_data)])

      ticket = Events::TicketPaid.new(:at => timestamp, :ticketID => @ticketID)
      encoded_data = Events::TicketPaid.encode(ticket)
      res = $pg.exec_params("INSERT INTO events (type, data) VALUES ($1,$2)", ["TicketPaid", $pg.escape_bytea(encoded_data)])

      ticket = Events::TicketComplete.new(:at => timestamp, :ticketID => @ticketID)
      encoded_data = Events::TicketComplete.encode(ticket)
      res = $pg.exec_params("INSERT INTO events (type, data) VALUES ($1,$2)", ["TicketComplete", $pg.escape_bytea(encoded_data)])

      @requestSent = Time.now.to_i
      @response = $client.post "#{$endpoint}/ticket/#{@ticketID}/complete"
    end

    it "returns 404 NOT FOUND" do
      expect(@response.status).to eq 404
    end
  end

  context "unknown ticket" do
    before :all do
      @ticketID = SecureRandom.uuid
      @response = $client.post "#{$endpoint}/ticket/#{@ticketID}/complete"
    end

    it "returns 404 NOT FOUND" do
      expect(@response.status).to eq 404
    end
  end
end
