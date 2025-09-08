# Project Improvement TODO List

## High-Level Improvements

-   [ ] **Add a `README.md` file:** The project is missing a `README.md` file. This file should explain what the project does, how to install it, and how to use it.
-   [ ] **Add Tests:** There are no tests in the project. Adding unit tests for the utility functions and command logic would improve the code quality and prevent regressions.
-   [ ] **Add a License:** The `LICENSE` file is empty. It would be good to add a license to the project, such as MIT or Apache 2.0.
-   [ ] **Improve Error Handling:** The project should not use `log.Fatal` to handle errors. Instead, it should return errors and handle them gracefully, printing user-friendly messages.
-   [ ] **Improve Configuration Management:** The configuration management could be improved by adding validation and better default values.
-   [ ] **Add a `version` command:** It's a good practice to have a `version` command that prints the version of the CLI.
-   [ ] **Improve Code Style and Consistency:** The code style should be consistent throughout the project. For example, all the error messages should be in the same language.
-   [ ] **Improve Modularity:** Some of the logic in the `cmd` files could be moved to the `utils` package to improve modularity and reusability.
-   [ ] **Use a structured logger:** The project uses `fmt.Println` and `log.Fatal` for logging. It would be better to use a structured logger like `logrus` or `zap`.

## `cmd` Package Improvements

### `blog` command

-   [ ] Implement the functionality to create a new blog post.
-   [ ] Update the `help` text to reflect the actual functionality.

### `lab` command

-   [ ] Replace `log.Fatal` with proper error handling.
-   [ ] Load templates from files instead of hardcoding them.
-   [ ] Extend language support beyond C and C++.
-   [ ] Use correct file permissions (`0644` for files, `0755` for directories).
-   [ ] Avoid using `os.Chdir`. Work with absolute paths.
-   [ ] Check if the editor is installed before trying to open it.

### `list` command

-   [ ] Refactor the `Run` function to reduce code duplication.
-   [ ] Improve the TUI with more features (e.g., opening the project, deleting a project).
-   [ ] Add more filtering options (e.g., by date, by name).

### `set` command

-   [ ] Add validation for input values (e.g., year, semester).
-   [ ] Use an empty string as the default value for `labs_location` instead of "DEFAULT".
-   [ ] Replace `log.Fatal` with proper error handling.

## `utils` Package Improvements

### `dirCrawler.go`

-   [ ] Replace `log.Fatal` with proper error handling.
-   [ ] Use `filepath.Join` for path construction.

### `markdown.go`

-   [ ] Return errors from `ParseMdString` instead of just logging them.

### `projectManager.go`

-   [ ] Replace `log.Fatal` with proper error handling.
-   [ ] Make the fallback `labs_location` configurable.
-   [ ] Use correct file permissions (`0755`) in `CreateDirectory`.
-   [ ] Check if a directory already exists in `CreateDirectory`.
-   [ ] Refactor `GetProjectsMetadata` to use the `DirCrawler` function.

### `tui.go`

-   [ ] Make the TUI styles configurable.
-   [ ] Add more features to the TUI (e.g., opening the project, deleting a project, showing the README).