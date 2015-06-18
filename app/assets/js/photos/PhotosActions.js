import alt from '../alt.js'

class PhotosActions {
    updatePhotos(photos) {
        this.dispatch(photos);
    }

    deletePhoto(id) {
        this.dispatch(id);
    }
}

module.exports = alt.createActions(PhotosActions);
