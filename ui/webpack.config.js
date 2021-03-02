const path = require('path');
const ESLintPlugin = require('eslint-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const TsconfigPathsPlugin = require('tsconfig-paths-webpack-plugin');

const isDev = process.env.NODE_ENV === 'development';

const rules = [
  {
    test: /\.(c|sc|sa)ss$/,
    exclude: /node_modules/,
    use: [
      MiniCssExtractPlugin.loader,
      { loader: 'css-loader' },
      {
        loader: 'postcss-loader',
        options: { sourceMap: isDev }
      },
      {
        loader: 'sass-loader',
        options: {
          sourceMap: isDev,
          implementation: require('sass')
        }
      }
    ]
  },
  {
    test: /\.ts(x?)$/,
    exclude: /node_modules/,
    use: ['ts-loader']
  },
  {
    test: /\.(woff(2)?|ttf|eot|svg)(\?v=\d+\.\d+\.\d+)?$/,
    use: [
      {
        loader: 'file-loader',
        options: {
          limit: 10000,
          name: 'fonts/[name].[ext]',
          mimetype: 'application/font-woff'
        }
      }
    ]
  }
];

module.exports = {
  entry: {
    main: './src/js/index.tsx'
  },
  name: 'main',
  mode: isDev ? 'development' : 'production',
  devtool: isDev ? 'source-map' : false,
  output: {
    path: path.resolve(__dirname, 'dist'),
    publicPath: '/dist/',
    filename: isDev ? '[name].js' : '[name].[hash].js',
    chunkFilename: isDev ? '[name].chunk.js' : '[name].chunk.[hash].js'
  },
  module: { rules },
  plugins: [
    new MiniCssExtractPlugin({
      filename: !isDev ? '[name].[contenthash].css' : '[name].css'
    }),
    new ESLintPlugin({
      extensions: ['ts', 'tsx'],
    }),
  ].filter(Boolean),

  resolve: {
    extensions: ['.ts', '.tsx', '.js', '.jsx'],
    plugins: [new TsconfigPathsPlugin()]
  }
};
