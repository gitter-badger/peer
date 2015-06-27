import React from 'react'

import API from './API.js'
import PhotosActions from './photos/PhotosActions.js'

export default
class GridPhoto extends React.Component {

    constructor(props) {
        super(props);

        this.handlePhotoDelete = this.handlePhotoDelete.bind(this);
    }

    render() {
        var downloadPath = "/photos/" + this.props.photo.id + "/thumbnail";
        return (
            <div className="grid-photo" key={this.props.photo.id}>
                <img src={downloadPath} alt={this.props.photo.file_name} height="220"/>
                <div className="overlay">
                    <div className="actions">
                        <i className="material-icons"
                           onTouchTap={this.handlePhotoDelete}>delete</i>
                    </div>
                </div>
            </div>
        )
    }

    handlePhotoDelete() {
        API.deletePhoto(this.props.photo.id)
            .then((response) => {
                PhotosActions.deletePhoto(this.props.photo.id);
            })
    }
}
