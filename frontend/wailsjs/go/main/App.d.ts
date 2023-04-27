// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';
import {backend} from '../models';

export function AddNewSave(arg1:string,arg2:string,arg3:number):Promise<number>;

export function AddNewSeason(arg1:number,arg2:main.NewSeason):Promise<void>;

export function GetImage(arg1:string):Promise<string>;

export function Greet(arg1:string):Promise<string>;

export function RetrieveSaves():Promise<Array<backend.Save>>;

export function SelectFileParse(arg1:string):Promise<string>;

export function SelectScheduleFile():Promise<string>;

export function SelectSquadFile():Promise<string>;

export function SelectTransfersFile():Promise<string>;

export function SingleImage(arg1:number):Promise<string>;

export function SingleSave(arg1:number):Promise<backend.Save>;

export function UploadSaveImage(arg1:number):Promise<string>;
