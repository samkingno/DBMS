import axios from 'axios';

const getDbinstances = () => axios.get('http://100.80.129.61:3000/api/v1/dbinstances').then(res => res.data);

const getDutys = () => axios.get('http://100.80.129.61:3000/api/v1/duty').then(res => res.data);

const createDuty = form => axios.post('http://100.80.129.61:3000/api/v1/duty',form).then(res=>res.data);

const postrequest = form => axios.post('http://100.80.129.61:3000/api/v1/pg_request',form).then(res=>res.data);

export {
    getDbinstances,
    getDutys,
    createDuty,
    postrequest,
};