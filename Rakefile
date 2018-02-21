require 'colorize'
require "rake/clean"
require "rspec/core/rake_task"

serviceName = "carpark"

task default: %w[all]

desc 'Run the unit tests'
task :unittest do
  puts "\nRake: Unit tests ...".colorize(:cyan)
  sh "go test ./..."
end

desc 'Build a local docker image'
task :build do
  puts "\nRake: Getting version ...".colorize(:cyan)
  semver = `curl -sL https://gist.githubusercontent.com/nicklanng/4e54bf35c13bf220408875dd059cad25/raw/44d53f8c1e58a2f73dcdb6eab36f0f6cefb6acc6/semver.sh | bash`.strip
  puts "Version is #{semver}".colorize(:green)

  puts "\nRake: Building Linux AMD64 ...".colorize(:cyan)
  ENV['GOOS'] = 'linux'
  ENV['GOARCH'] = 'amd64'
  ENV['CGO_ENABLED'] = '0'
  sh "go build -ldflags '-X main.version=#{semver}' -o bin/#{serviceName}-$GOOS-$GOARCH main.go"

  puts "\nRake: Building Docker image ...".colorize(:cyan)
  sh "docker build -t nicklanng/#{serviceName} ."
end

RSpec::Core::RakeTask.new(:spec) do |t|
  puts "\nRake: Verifying specifications ...".colorize(:cyan)
  t.pattern = Dir.glob("specs/**/*.rb")
  t.rspec_opts = "--format documentation"
end
task :spec

task :all => [:unittest, :build, :spec]

CLEAN << "bin"
