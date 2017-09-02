const fs = require('fs');
const autobahn = require('autobahn');
const commandLineArgs = require('command-line-args');
 
const optionDefinitions = [
  { name: 'host', alias: 'h', type: String, defaultValue: 'localhost' },
  { name: 'port', alias: 'p', type: Number, defaultValue: 8000 },
  { name: 'secure', alias: 's', type: Boolean, defaultValue: false },
  { name: 'realm', alias: 'r', type: String, defaultValue: 'realm1' },
  { name: 'file', alias: 'f', type: String, defaultValue: 'ticks.json' },
  { name: 'delay', alias: 'd', type: Number, defaultValue: 100 }
];

const options = commandLineArgs(optionDefinitions);

const connection = new autobahn.Connection({
  url: (options.secure ? 'wss' : 'ws') + '://' + options.host + ':' + options.port,
  realm: options.realm
});

const ticksFile = options.file;
let ticks = [];

const publishDelayMs = options.delay;

fs.exists(ticksFile, function(exists) {
  if (exists) {
    fs.readFile(ticksFile, function(err, data) {
      if (err) {
        console.error('Error:', err);
      } else {
        ticks = JSON.parse(data);

        connect();
      }
    });
  }
});

function publish(session) {
  const tick = ticks.pop();

  if (tick) {
    console.log("Publishing to 'ticker' topic, " + ticks.length + " ticks remaining");

    session.publish('ticker', tick);

    setTimeout(() => publish(session), publishDelayMs);
  } else {
    console.log('No more ticks, exiting.');

    process.exit();
  }
}

function connect() {
  console.log('Ticks:', ticks.length);

  connection.onopen = (session) => {
    console.log('Websocket connect opened!');

    publish(session);
  };

  connection.onclose = () => {
    console.log('Websocket connection closed.');
  };

  connection.open();
}

