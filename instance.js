import axios from "axios";
/*
export const createinstance = (addHeader) => {
    const Header = {
        "Content-Type": "application/json"
    }
    if (addHeader) {
        Header["Authorization"] = localStorage.getItem("token")
    }
    return axios.create({
        baseURL: 'http://localhost:8080/',
        timeout: 1000,
        headers: Header
    });
}*/

export const publicService = axios.create({
    baseURL: 'http://localhost:8080/',
    timeout: 1000,
    headers: {
        "Content-Type": "application/json"
    },
})
export const adminService = axios.create({
    baseURL: 'http://localhost:8080/',
    timeout: 1000,
    headers: {
        "Content-Type": "application/json",
        "Authorization": localStorage.getItem("token")
    },
})

