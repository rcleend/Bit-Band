const VueLoaderPlugin = require('vue-loader/lib/plugin')
const ExtactTextPlugin = require('extract-text-webpack-plugin')

const ExtractSASS = new ExtactTextPlugin('css/styles.css')

module.exports = ( env, options ) => {
	return {
		"mode": "production",
		"entry": "./src/index.js",
		"output": {
			"path": __dirname+'/public/assets/',
			"filename": "js/app.js"
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
					"use": ExtractSASS.extract([
						"css-loader",
					 	"sass-loader"
					])  
				}
			]
		},
		plugins: [
			new VueLoaderPlugin(),
			ExtractSASS
		]
	}
}	
