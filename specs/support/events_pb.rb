# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: events.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "events.TicketIssued" do
    optional :at, :message, 1, "google.protobuf.Timestamp"
    optional :ticketID, :string, 2
  end
  add_message "events.TicketPaid" do
    optional :at, :message, 1, "google.protobuf.Timestamp"
    optional :ticketID, :string, 2
  end
  add_message "events.TicketComplete" do
    optional :at, :message, 1, "google.protobuf.Timestamp"
    optional :ticketID, :string, 2
  end
end

module Events
  TicketIssued = Google::Protobuf::DescriptorPool.generated_pool.lookup("events.TicketIssued").msgclass
  TicketPaid = Google::Protobuf::DescriptorPool.generated_pool.lookup("events.TicketPaid").msgclass
  TicketComplete = Google::Protobuf::DescriptorPool.generated_pool.lookup("events.TicketComplete").msgclass
end
