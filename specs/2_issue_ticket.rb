require 'rspec'
require 'httpclient'
require 'json'
require 'google/protobuf'

describe "POST /ticket :: Issue a new ticket" do
  before :all do
    @requestSent = Time.now.to_i
    @response = $client.post "#{$endpoint}/ticket"
    sleep 5
  end

  it "returns 200 OK" do
    expect(@response.status).to eq 200
  end

  it "returns the ticket ID" do
    jsonResponse = JSON.parse @response.body
    isUUID = (jsonResponse["id"] =~ /^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i) == 0
    expect(isUUID).to be true
  end

  it "returns the time the ticket was issued" do
    jsonResponse = JSON.parse @response.body
    timeDelta = Time.parse(jsonResponse["issuedAt"]).to_i - @requestSent
    expect(timeDelta.abs).to be < 10
  end

  it "persists event to database" do
    rs = $pg.exec 'SELECT * FROM events WHERE seq = 1'

    seq = rs.getvalue 0, 0
    expect(seq).to eq "1"

    eventType = rs.getvalue 0, 1
    expect(eventType).to eq "TicketIssued"

    jsonResponse = JSON.parse @response.body

    binaryAsText = rs.getvalue 0, 2
    binary = $pg.unescape_bytea binaryAsText
    event = Events::TicketIssued.decode binary
    expect(event.ticketID).to eq jsonResponse["id"]

    timeDelta = event.at.seconds - @requestSent
    expect(timeDelta.abs).to be < 10
  end
end
