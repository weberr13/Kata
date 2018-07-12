#!/usr/bin/env ruby

class MegaGreeter
    attr_accessor :names

    # constructor
    def initialize(names = "World")
        @names = names
    end

    # say hi to everyone
    def say_hi
        if @names.nil?
            puts "..."
        elsif @names.respond_to?("each")
            @names.each do |name|
                puts "Hello #{name}"
            end
        else
            puts "Hello #{@names}"
        end
    end

    # say goodbye to everyone
    def say_bye
        if @names.nil?
            puts "..."
        elsif @names.respond_to?("join")
            puts "Goodbye #{names.join(", ")}"
        else
            puts "Goodbye #{names}"
        end
    end
end

if __FILE__ == $0
    mg = MegaGreeter.new
    mg.say_hi
    mg.say_bye

    mg.names = "bob"
    mg.say_hi
    mg.say_bye

    mg.names = [
        "bob",
        "terry"
    ]
    mg.say_hi
    mg.say_bye

    mg.names = nil
    mg.say_hi
    mg.say_bye
end