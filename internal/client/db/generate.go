package db

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../../bin/minimock -i Client -o ./mocks/ -s "_minimock.go"
//go:generate ../../../bin/minimock -i TxManager -o ./mocks/ -s "_minimock.go"
//go:generate ../../../bin/minimock -i Transactor -o ./mocks/ -s "_minimock.go"
//go:generate ../../../bin/minimock -i SQLExecer -o ./mocks/ -s "_minimock.go"
//go:generate ../../../bin/minimock -i NamedExecer  -o ./mocks/ -s "_minimock.go"
//go:generate ../../../bin/minimock -i QueryExecer -o ./mocks/ -s "_minimock.go"
//go:generate ../../../bin/minimock -i Pinger -o ./mocks/ -s "_minimock.go"
//go:generate ../../../bin/minimock -i DB -o ./mocks/ -s "_minimock.go"
