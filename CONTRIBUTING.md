# Contributing to Gocrafter

Thank you for considering contributing to `gocrafter`  
Please take a moment to review these guidelines to help us maintain consistency and quality across the project.

---

## Steps

1. Fork this repo
2. Create your feature branch (`git checkout -b feature/xyz`)
3. Commit your changes (`git commit -m 'Add xyz'`)
4. Push to the branch (`git push origin feature/xyz`)
5. Create a pull request

---

## Rules & Guidelines

### Branch Protection
- **Never push directly to `main` or `master`**.
- Always create a feature/fix branch (`feature/xyz`, `fix/typo`, etc.).
- Open a Pull Request (PR) for all changes.

### Pull Request Requirements
- All PRs must have a **clear title** and **concise description**.
- Include a short changelog at the top of the PR description in the format:


- Link related issues using `Closes #123` when applicable.

### Commit Message Format
- Use clear, conventional commits. Format:
    - [type]: short description

- Example:
    - **fix**: prevent banner from printing on subcommands
    - **feat**: add goose migration setup
    - **docs**: update usage section in README

---

### Required Documentation
- If your PR adds a new feature or significant fix, update the **`README.md`** with usage instructions.
- Include any **code examples or usage flows** in your PR description if applicable.

### Style & Lint
- Use `gofmt` and `go vet` before committing.
- Keep your code idiomatic and readable.

### Working with Templates
- When adding new `.tmpl` files:
  - Ensure they are embedded correctly.
  - Keep logic declarative and minimal inside templates.
  - Validate generated output with `gocrafter init`.

---

## Additional Notes

- Contributions are reviewed within a few days.
- Feel free to discuss ideas by opening an issue before coding.
- Respect reviewer feedback and collaborate professionally.

---

Thanks again for contributing to `gocrafter`!