import { Record } from "./record.model";

export class Certificate {


    public key: number;
    public record: Record;


    constructor(data: any) {
        this.key = data.Key ? data.Key : null;
        this.record = data.Record ? new Record(data.Record) : new Record({});



    }
}