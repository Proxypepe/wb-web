import axios, {AxiosResponse} from "axios";
import {IOrder} from "../models/Order";

export default class OrderService {
    static async getOrderByUid(uid: string): Promise<AxiosResponse<IOrder>> {
        return axios.get<IOrder>(`http://localhost:8080/order?order_uid=${uid}`)
    }
}