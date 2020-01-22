# Zach Thomas project-0
This project will allow a user to see various player individual stats in the NBA. Data collected January 8, 2020 from https://www.basketball-reference.com.

This project also contains a dataframe package that allows for handling of csv data in a convenient way.

There are two executable files:

statsCLI.go is a command line interface where a user passes in either the flag "-player" followed by a player name to retrieve that player's averge statistics or a user passes in the flag "-stat" followed by the statistic they are interested in and to the program will print to the console all players' statistics sorted by the input stat.

statsd.go runs a more user friendly version of the program on localhost:8080.

## Desired Functionality
- [x] input player name receive stats
- [X] input specific stat and receive leaders in that stat
- [ ] input player and stat to receive specific data
- [ ] retrieve updated stats