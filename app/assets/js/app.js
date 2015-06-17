import React from 'react'
import {AppBar, FloatingActionButton} from 'material-ui'

require('es6-promise').polyfill();
require('isomorphic-fetch');

var injectTapEventPlugin = require("react-tap-event-plugin");
injectTapEventPlugin();

var mui = require('material-ui');
var ThemeManager = new mui.Styles.ThemeManager();

var API = {
    photos() {
        return fetch('/photos')
            .then(this._status)
            .then(this._json)
    },

    upload(file) {
        var form = new FormData()
        form.append('file', file)

        return fetch('/photos', {
            method: 'post',
            body: form
        })
    },

    deletePhoto(id) {
        return fetch('/photos/' + id, {
                method: 'delete',
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
        }

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
                    iconClassName="muidocs-icon-content-add"
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

class GridPhoto extends React.Component {

    constructor(props) {
        super(props)

        this.handlePhotoDelete = this.handlePhotoDelete.bind(this);
    }

    render() {
        var downloadPath = "/photos/" + this.props.photo.id + "/download";
        return (
            <div key={this.props.photo.id}
                 style={{display: 'inline-block', marginRight: '10px', position: 'relative'}}>
                <i className="material-icons"
                    onTouchTap={this.handlePhotoDelete}
                    style={{position: 'absolute', top: 5, right: 5, color: '#eee'}}>delete</i>
                <img src={downloadPath} alt={this.props.photo.file_name} height="220"/>
            </div>
        )
    }

    handlePhotoDelete() {
        API.deletePhoto(this.props.photo.id)
            .then((response) => {
                console.log(response);
            })
    }
}

App.childContextTypes = {
    muiTheme: React.PropTypes.object
};

React.render(<App />, document.getElementById('app'));
