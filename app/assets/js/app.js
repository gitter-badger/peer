import React from 'react'

class Test extends React.Component {
    render() {
        return (
            <div>Hello World, this is React.</div>
        )
    }
}

React.render(<Test />, document.getElementById('app'));
