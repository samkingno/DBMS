import axios from 'axios';

const getDbinstances = () => axios.get('http://localhost:3000/api/v1/dbinstances').then(res => res.data);

export {
    getDbinstances,
};