# Quiz

## Goal
- Learn how to read file
- Learn hot to read .csv file
- Learn how to use Flag

## Note
[how to read file](https://stackoverflow.com/questions/36111777/golang-how-to-read-a-text-file)
### Open file
```go
import (
	"os"
	"io/ioutil"
)

file, err := os.Open("problem.csv")
if err != nil {
	log.Fatal(err)
}

defer file.Close()
```

### Read all
```go
content, err := ioutil.ReadAll(file)
fmt.Println(string(file))
```

### Read line
```go
scanner := bufio.NewScanner(file)

for scanner.Scan() {
	fmt.Println(scanner.Text())
	fmt.Println(scanner.Bytes())
}
```

### buffer-size
```go
buf := make([]byte, 5)

for {
	n, err := file.Read(buf)

	if n > 0 {
		fmt.Println(string(buf[:n]))
	}

	if err == io.EOF {
		break
	}

	if err != nil {
		log.Printf("read %d bytes: %v", n, err)
		break
	}
}
```

### Read .csv file
```go
lines, err := csv.NewReader(file).ReadAll()
if err != nil {
	log.Fatal(err)
}
```