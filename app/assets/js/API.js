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

export default API;
