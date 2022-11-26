export interface File {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt?: any;
    name: string;
    location: string;
    type: string;
    size: number;

    selected: boolean;
}