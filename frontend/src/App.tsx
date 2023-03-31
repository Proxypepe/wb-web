import React, {useState} from 'react';
import type { FC } from 'react';
import {Divider, Input, Layout} from 'antd';
import 'antd/dist/reset.css';
import './App.css';
import OrderService from "./api/OrderService";
import {IOrder} from "./models/Order";
import {OrderCard} from "./components/OrderCard";

const { Search } = Input;

const App: FC = () => {

    const [loading, setLoading] = useState(false);
    const [order, setOrder] = useState<IOrder>();
    const [reqResult, setReqResult] = useState("Make request");

    const onSearch = (value: string) => {
        setLoading(true)
        setReqResult("Making request...")
        try{
            setTimeout(( async () => {
                    await OrderService.getOrderByUid(value)
                        .then((response) => {
                            if (response.status === 200) {
                                console.log(response.data)
                                setOrder(response.data)
                                setLoading(false)
                                return
                            }
                        })
                        .catch( e => {
                            console.log(e)
                            setReqResult("Nothing was found")
                            setLoading(false)
                            setOrder(undefined)
                        })
                }
            ), 1000 )
        } catch (e) {
            setLoading(false)
        }
    };

    return (
        <div className="App">
            <Layout>
                <Layout.Content style = {{padding: 10, paddingTop: 20, background: 'white'}}>
                    <Search placeholder="Enter order uid..." style={{ width: "40%" }} loading={loading} onSearch={onSearch}/>
                    <Divider orientation="left">Request result</Divider>
                    <OrderCard order={order} reqResult={reqResult}/>
                </Layout.Content>
            </Layout>
        </div>
        )
};

export default App;