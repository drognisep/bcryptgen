# BCrypt Generator
This tool allows the user to generate passwords, generate Bcrypt hashes, and check user provided passwords and/or hashes.

## Installation
Download the latest release package for your operating system and place it somewhere accessible. The executable is self-contained, so there is no explicit install step needed.

If you'd rather build it yourself, then clone down the repository and make sure you have everything required to build the app on your OS by checking the [fyne getting started docs](https://fyne.io/develop/).

## What is BCrypt?
[BCrypt](https://auth0.com/blog/hashing-in-action-understanding-bcrypt/) is a variable complexity hashing scheme used to protect sensitive information - like passwords - from being easily discovered by approaches like comparing against [rainbow tables](https://en.wikipedia.org/wiki/Rainbow_table). Since the hashing process is a one-way operation, the hash can be stored in the clear in a file or database without fear of accidentally exposing the original value.

**Note:** The original password *cannot* be recovered from the BCrypt hash. A password *match* can be determined by hashing an input value using the same salt as the stored hash and comparing them together.

## Password Security
While BCrypt makes password storage significantly more secure, the underlying password should be secure as well. This application uses  the [go-password](https://github.com/sethvargo/go-password) library, which relies on a cryptographically secure pseudo-random number generator ([CSPRNG](https://en.wikipedia.org/wiki/Cryptographically_secure_pseudorandom_number_generator)).

The generator can adhere to different password policies by simply changing the generator settings in the app, but reasonable defaults will be assigned when the user changes the password length.
