export class Record {
    public collegeName: string;
    public examination: string;
    public prNo: number;
    public seatNo: number;
    public studentName: string;
    public sub: string;
    public yearOfPassing: number;

    constructor(data: any) {
        this.collegeName = data.College_Name ? data.College_Name : null;
        this.examination = data.Examination ? data.Examination : null;
        this.prNo = data.PR_no ? data.PR_no : null;
        this.seatNo = data.Seat_no ? data.Seat_no : null;
        this.studentName = data.Student_Name ? data.Student_Name : null;
        this.sub = data.Sub ? data.Sub : null;
        this.yearOfPassing = data.Year_Of_Passing ? data.Year_Of_Passing : null;
    }
}