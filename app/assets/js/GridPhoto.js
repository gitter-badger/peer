import React from 'react'

export default
class GridPhoto extends React.Component {

    constructor(props) {
        super(props);

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
