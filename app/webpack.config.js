module.exports = ( env, options ) => {
	return {
		"mode": "development",
		"entry": "./src/index.js",
		"output": {
			"path": __dirname+'/public/assets/js',
			"filename": "app.js"
		},
		"module": {
			"rules": [
				{
					"test": /\.js$/,
					"exclude": /node_modules/,
					"use": {
						"loader": "babel-loader",
						"options": {
							"presets": [
								"env"
							]
						}
					}
				},
				{
					"test": /\.scss$/,
					"use": [
						"style-loader",
						"css-loader",
						"sass-loader"
					]
				}
			]
		}
	}
}	
