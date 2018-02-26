require "httpclient"
require "docker"
require 'json'
require 'pg'

RSpec.configure do |config|
  postgresContainer = nil
  $carparkContainer = nil

  config.before(:suite) do
    # start postgres server
    postgresContainer = create_postgres
    postgresContainer.start
    postgresPort = postgresContainer.json["NetworkSettings"]["Ports"]["5432/tcp"][0]["HostPort"]
    postgresName = postgresContainer.json["Name"]

    sleep 10

    # create DATABASE
    $pg = PG.connect(
      :host => 'localhost',
      :port => postgresPort,
      :dbname => 'postgres',
      :user => 'postgres',
      :password => 'postgres',
    )
    $pg.exec("CREATE DATABASE carpark")
    $pg.close()
    $pg = PG.connect(
      :host => 'localhost',
      :port => postgresPort,
      :dbname => 'carpark',
      :user => 'postgres',
      :password => 'postgres',
    )

    # start carpark server
    $carparkContainer = create_carpark(postgresName)
    $carparkContainer.start
    port = $carparkContainer.json["NetworkSettings"]["Ports"]["8443/tcp"][0]["HostPort"]
    $endpoint = "https://127.0.0.1:#{port}"

    # create http client
    $client = HTTPClient.new
    $client.ssl_config.verify_mode = OpenSSL::SSL::VERIFY_NONE

    # wait for servers to start
    sleep 2
  end

  config.after(:suite) do
    $carparkContainer.stop
    postgresContainer.stop
  end
end

def create_carpark(postgresName)
  return Docker::Container.create(
    "Image" => "nicklanng/carpark:dev",
    "Env" => [
      "DB_USER=postgres",
      "DB_PASSWORD=postgres",
      "DB_HOST=postgres",
      "CERT_PATH=/support/server.crt",
      "KEY_PATH=/support/server.key"
    ],
    "HostConfig" => {
      "PortBindings" => {
        "8443/tcp" => [{}]
      },
      "Binds" => [
        "#{File.join(Dir.pwd, "specs", "support")}:/support",
      ],
      'Links' => [
        "#{postgresName}:postgres",
      ]
    }
  )
end

def create_postgres()
  puts 'Getting Postgres Image...'
  Docker::Image.create('fromImage' => 'postgres:10.2')
  return Docker::Container.create(
    "Image" => "postgres:10.2",
    "Env" => [
      "POSTGRES_PASSWORD=postgres"
    ],
    "HostConfig" => {
      "PortBindings" => {
        "5432/tcp" => [{}]
      }
    }
  )
end
