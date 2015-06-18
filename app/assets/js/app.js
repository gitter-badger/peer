require('es6-promise').polyfill();
require('isomorphic-fetch');
var injectTapEventPlugin = require("react-tap-event-plugin");
injectTapEventPlugin();

var mui = require('material-ui');
var ThemeManager = new mui.Styles.ThemeManager();

import React from 'react'
import {AppBar, FloatingActionButton} from 'material-ui'
import GridPhoto from './GridPhoto.js'

var API = {
    photos() {
        return fetch('/photos')
            .then(this._status)
            .then(this._json)
    },

    upload(file) {
        var form = new FormData();
        form.append('file', file);

        return fetch('/photos', {
            method: 'post',
            body: form
        })
    },

    deletePhoto(id) {
        return fetch('/photos/' + id, {
            method: 'delete'
        })
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
        };

        this.inputFileChange = this.inputFileChange.bind(this);
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
        var renderGridPhoto = function (photo, index) {
            return (
                <GridPhoto photo={photo}/>
            )
        };

        return (
            <div>
                <AppBar
                    title="Photos"
                    showMenuIconButton={false}/>

                <div style={{margin: '24px'}}>
                    <div>
                        <input
                            type="file"
                            onChange={this.inputFileChange}/>
                    </div>
                    {this.state.photos.map(renderGridPhoto)}
                </div>
                <FloatingActionButton
                    style={{position: 'fixed', bottom: '20px', right: '20px'}}>
                    <i className="material-icons" style={{color: 'white'}}>cloud_upload</i>
                </FloatingActionButton>
            </div>
        )
    }

    inputFileChange(event) {
        API.upload(event.target.files[0]);
    }
}

App.childContextTypes = {
    muiTheme: React.PropTypes.object
};

React.render(<App />, document.getElementById('app'));
