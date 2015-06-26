import alt from '../alt.js'
import _ from 'lodash'
import PhotosActions from './PhotosActions.js'

class PhotosStore {
    constructor() {
        this.photos = [];

        this.bindListeners({
            handleUpdatePhotos: PhotosActions.UPDATE_PHOTOS,
            handleUploadedPhoto: PhotosActions.UPLOADED_PHOTO,
            handleDeletePhoto: PhotosActions.DELETE_PHOTO
        });
    }

    handleUpdatePhotos(photos) {
        this.photos = photos;
    }

    handleUploadedPhoto(photo) {
        console.log(photo);
        this.photos.push(photo);
    }

    handleDeletePhoto(id) {
        _.remove(this.photos, (photo) => {
            return photo.id == id;
        });
    }
}

module.exports = alt.createStore(PhotosStore, 'PhotosStore');
