const rules = require('./webpack.rules');

rules.push({
  test: /\.css$/,
  use: [
    { loader: 'style-loader' },
    { loader: 'css-loader' },
    { loader: 'postcss-loader' } // v5.0.0 requires webpack 5. See inside the electron forge plugin
  ],
});

module.exports = {
  // Put your normal webpack config below here
  module: {
    rules,
  },
};
