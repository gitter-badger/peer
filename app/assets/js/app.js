import React from 'react'
import {AppBar} from 'material-ui'

require('es6-promise').polyfill();
require('isomorphic-fetch');

var mui = require('material-ui');
var ThemeManager = new mui.Styles.ThemeManager();

var API = {
    photos() {
        return fetch('/photos')
            .then(this._status)
            .then(this._json)
    },

    _status(response) {
        if (response.status >= 200 && response.status < 300) {
            return response;
        }
        throw new Error(response.statusText);
    },

    _json(response) {
        return response.json()
    }
};

class App extends React.Component {

    constructor() {
        super();

        this.state = {
            photos: []
        }
    }

    getChildContext() {
        return {
            muiTheme: ThemeManager.getCurrentTheme()
        };
    }

    componentDidMount() {
        API.photos()
            .then((photos) => {
                this.setState({
                    photos: photos
                })
            })
    }

    render() {
        var renderPhoto = function (photo, index) {
            return (
                <li key={index}>{photo.id} - {photo.file_name}</li>
            )
        };

        return (
            <div>
                <AppBar
                    title="Photos"
                    showMenuIconButton={false}/>
                <ul>{this.state.photos.map(renderPhoto)}</ul>
            </div>
        )
    }
}

App.childContextTypes = {
    muiTheme: React.PropTypes.object
};

React.render(<App />, document.getElementById('app'));
