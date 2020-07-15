import axios from 'axios'


export default function apiCheck(){
    axios.get('http://127.0.0.1:8080/api/get')
    .then(res => {
        console.log(res)
    })
}