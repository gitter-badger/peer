import alt from '../alt.js'

class PhotosActions {
    updatePhotos(photos) {
        this.dispatch(photos);
    }

    uploadedPhoto(photo) {
        this.dispatch(photo);
    }

    deletePhoto(id) {
        this.dispatch(id);
    }
}

module.exports = alt.createActions(PhotosActions);
