export interface IOrder {
    order_uid:          string;
    track_number:       string;
    entry:              string;
    delivery:           IDelivery;
    payment:            IPayment;
    items:              IItem[];
    locale:             string;
    internal_signature: string;
    customer_id:        string;
    delivery_service:   string;
    shardkey:           string;
    sm_id:              number;
    date_created:       string;
    oof_shard:          string;
}

export interface IDelivery {
    name:    string;
    phone:   string;
    zip:     string;
    city:    string;
    address: string;
    region:  string;
    email:   string;
}

export interface IItem {
    chrt_id:      number;
    track_number: string;
    price:        number;
    rid:          string;
    name:         string;
    sale:         number;
    size:         string;
    total_price:  number;
    nm_id:        number;
    brand:        string;
    status:       number;
}

export interface IPayment {
    transaction:   string;
    request_id:    string;
    currency:      string;
    provider:      string;
    amount:        number;
    payment_dt:    number;
    bank:          string;
    delivery_cost: number;
    goods_total:   number;
    custom_fee:    number;
}
