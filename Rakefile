require 'colorize'
require "rake/clean"
require "rspec/core/rake_task"

serviceName = "carpark"

task default: %w[all]

desc 'Download required tools'
task :install do
 sh("go get -u github.com/jteeuwen/go-bindata/...")
 sh("go install ./vendor/github.com/golang/protobuf/protoc-gen-go")
 sh("go get -u -a golang.org/x/tools/cmd/stringer")
end

desc 'Create generated code'
task :codegen do
  puts "\nGenerating Go code ...".colorize(:cyan)
 sh("go generate ./...")
end

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
  sh "docker build -t nicklanng/#{serviceName}:dev ."
end

task :rubycodegen do
  puts "\nGenerating Ruby code ...".colorize(:cyan)
  sh("protoc -I ./events --ruby_out ./specs/support ./events/events.proto")
end

RSpec::Core::RakeTask.new(:spec) do |t|
  puts "\nRake: Verifying specifications ...".colorize(:cyan)
  t.pattern = Dir.glob("specs/**/*.rb")
  t.rspec_opts = "--format documentation"
end
task :spec => :rubycodegen

task :all => [:codegen, :unittest, :build, :spec]

CLEAN << "bin"
