# Task Scheduler

The Task Scheduler is a Go application that allows you to efficiently distribute and process tasks using a pool of workers. It leverages concurrency, channels, and goroutines to achieve high performance and parallel processing.

## Features

- Distributes tasks among multiple workers for concurrent execution.
- Utilizes goroutines and channels to enable efficient communication and coordination between workers.
- Provides a flexible and scalable solution for handling large sets of tasks.

## Getting Started

### Prerequisites

- Go programming language (version 1.19 or higher)
- [Install Go](https://golang.org/doc/install)

### Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/ivofreitas/task-scheduler-sample.git
   ```

2. Change into the project directory:

   ```shell
   cd task-scheduler-sample
   ```

3. Build the project:

   ```shell
   go build
   ```

### Usage

1. Modify the `main.go` file to customize the list of tasks and the number of workers according to your requirements.

2. Run the application:

   ```shell
   ./task-scheduler-sample
   ```

3. View the results in the console output, which displays the completion status of each task.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to submit a pull request or open an issue in the GitHub repository.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

- [Go programming language](https://golang.org/)
- [OpenAI GPT-3](https://openai.com/)

## Contact

For any inquiries or questions, please contact ivorfn@gmail.com.

Feel free to modify the sections and content according to your project's specific details and requirements. The README file serves as a central place for users and contributors to understand how to use and contribute to your project, so make sure to provide clear instructions and relevant information.