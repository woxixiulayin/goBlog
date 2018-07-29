const path = require('path');
const Clean = require('clean-webpack-plugin');
const AssetsPlugin = require('assets-webpack-plugin');
const outputDir = path.join(__dirname, "dist/");
const isProd = process.env.NODE_ENV === 'production';
const webpack = require('webpack');

function resolve(dir) {
    return path.join(__dirname, './', dir)
}

const plugins = [
    new Clean(['dist'], { root: './'}),
    new AssetsPlugin({ filename: 'assets.json' }),
    new webpack.ProvidePlugin({
        React: 'react',
        log: 'src/lib/log',
    }),
]

const entry = {
    home: './src/index.jsx'
}

module.exports = {
  entry,
  mode: isProd ? 'production' : 'development',
  output: {
    path: outputDir,
    publicPath: '/assets/dist',
    filename: 'index.[hash].js',
  },
  resolve: {
      modules: [
          resolve('./node_modules')
      ],
      extensions: ['.js', '.jsx'],
      alias: {
          css: resolve("./css"),
          src: resolve('./src'),
      }
  },
  module: {
      rules: [
        { test:/\.scss$/,
        use: [
            { loader: 'style-loader' },
            { loader: 'css-loader' },
            {
              loader: 'sass-loader',
              options: {
                modules: true
              }
            }
          ]
        },
        { test: /\.js|jsx$/,
          use: [
            { loader: "babel-loader"}
            ],
          exclude: /node_modules/,
        }
      ]
  },
  plugins,
};
