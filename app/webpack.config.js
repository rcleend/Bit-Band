const VueLoaderPlugin = require('vue-loader/lib/plugin')

module.exports = ( env, options ) => {
	return {
		"mode": "production",
		"entry": "./src/index.js",
		"output": {
			"path": __dirname+'/public/assets/js',
			"filename": "app.js"
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
					"use": [
						"vue-style-loader",
						"css-loader",
						{
							"loader": "sass-loader",
							"options": {
								"data": `@import "./src/assets/scss/_global.scss";`
							}
						}
					]
				}
			]
		},
		plugins: [
			new VueLoaderPlugin()
		]
	}
}	
