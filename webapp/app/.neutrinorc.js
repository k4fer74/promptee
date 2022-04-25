const react = require('@neutrinojs/react');

module.exports = {
  options: {
    root: __dirname,
    mains: {
      index: {
        'entry': 'index',
        'title': 'Get started',
      },
      bible: {
        'entry': 'bible',
        'title': 'Bible',
      },
      prompter: {
        'entry': 'prompter',
        'title': 'Prompter',
      },
      lyrics: {
        'entry': 'lyrics',
        'title': 'Lyrics',
      }
    },
  },
  use: [
    react({
      meta: {
        viewport: 'width=device-width, initial-scale=1',
      },
    }),
  ],
};
