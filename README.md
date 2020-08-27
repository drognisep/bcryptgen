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

# How to use it
There are only 3 distinct steps to getting a BCrypt hash with this application.

## 1. Generate the password
When the app first loads the user will be presented with a blank form.

![Initial State](https://github.com/drognisep/bcryptgen/blob/master/docs/images/blank.png?raw=true)

Press the Generate Password button to get the ball rolling. Configure your password according to your needs and click on Generate. Alternatively, you can enter your own password without going through the generation step. The longer the password, the more secure it is.

![Generate Password](https://github.com/drognisep/bcryptgen/blob/master/docs/images/genpass.png?raw=true)

**Note:** The application can generate up to a 64 character password.

## 2. Generate the hash
Once you have a password it's time to generate a hash. Just click on the Generate Hash button.

![Generate Hash](https://github.com/drognisep/bcryptgen/blob/master/docs/images/genhash.png?raw=true)

## 3. Check the hash and password
Just for good measure, click on the Compare password and hash button to confirm that everything is working as it should.

![Generate Hash](https://github.com/drognisep/bcryptgen/blob/master/docs/images/match.png?raw=true)

## All set!
If you experience any issues, please feel free to create an issue here so I can get it fixed as soon as possible. Thanks!
