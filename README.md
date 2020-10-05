# Diodon

Diodon is a discord bot that learn from your sentences to make his own.

## Presentation

This discord bot learn form your message and make sentences with a markov chain
algorithm.

## Usage

You have nothing to do to make him learn.  
He will answer you when you mention him.

## Envorinment variable (cf. `.env` file)

The `.env` file can be filled with the help of the `.env.tpl` file.

### Variables

The `DISCORD` variable is the discord token needed to connect to discord

The `TALKINESS` variable will define the frequency the bot answer by him self
to your messages. It should be a float. When a message is send a random
number in the interval `[0.0, 1.0[` will be choosen. If the choosen number is
lower to the `TALKINESS` the bot will spead (that mean higher is more
talkiness).

The `CONNECTION_STRING` variable is given to the
[markovchaingo](https://github.com/keftcha/markovchaingo) lib.  
Here the docker compose is made to work with the `file:///data.json` connection
string and it is a volume to the `data.json` file at the root of the project
(the `data.json` file isn't commited).  
Be pleased to fork and modify the bot for your need.
