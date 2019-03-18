const VueLoaderPlugin = require('vue-loader/lib/plugin');

// This Plugin is used to output seperate CSS files
const ExtractTextPlugin = require('extract-text-webpack-plugin');

// This Plugin is used to minify the outputted CSS files
const OptimizeCssAssetsPlugin = require('optimize-css-assets-webpack-plugin');

module.exports = ( env, options ) => {
	return {
		"mode": "development",
		"entry": "./src/index.js",
		"output": {
			"path": __dirname+'/dist/',
			"filename": "js/app.min.js"
		},
		"module": {
			"rules": [
				{
					"test": /\.vue$/,
					"loader": "vue-loader"
				},
				{
					"test": /\.js$/,
					"exclude": /node_modules/,
					"use": {
						"loader": "babel-loader",
						"options": {
							"presets": [
								"@babel/env"
							]
						}
					}
				},
				{
					"test": /\.scss$/,
					"use": ExtractTextPlugin.extract([
						"css-loader",
						'sass-loader',
					])	
				}
			]
		},
		plugins: [
			new VueLoaderPlugin(),
			new ExtractTextPlugin('css/styles.min.css'),
			new OptimizeCssAssetsPlugin()
		]
	}
}	
