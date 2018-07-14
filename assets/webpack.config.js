const path = require('path');
const Clean = require('clean-webpack-plugin');
const AssetsPlugin = require('assets-webpack-plugin');
const outputDir = path.join(__dirname, "build/");
const isProd = process.env.NODE_ENV === 'production';

const plugins = [
    new Clean(['dist'], { root: './'}),
    new AssetsPlugin({ filename: 'assets.json' }),
]

const entry = {
    home: './src/index.bs.js'
}

module.exports = {
  entry,
  mode: isProd ? 'production' : 'development',
  output: {
    path: outputDir,
    publicPath: '/assets/build',
    filename: 'index.[hash].js',
  },
  plugins,
};
