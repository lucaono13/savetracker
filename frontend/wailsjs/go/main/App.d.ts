// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';
import {backend} from '../models';

export function AddNewSave(arg1:string,arg2:string,arg3:number,arg4:string):Promise<main.ErrorReturn>;

export function AddNewSeason(arg1:number,arg2:main.NewSeason):Promise<main.ErrorReturn>;

export function DeleteSave(arg1:number):Promise<main.ErrorReturn>;

export function GetAllPlayersSeason():Promise<main.ErrorReturn>;

export function GetAllPlayersTotals():Promise<main.ErrorReturn>;

export function GetAllRankings():Promise<main.ErrorReturn>;

export function GetAllResults():Promise<main.ErrorReturn>;

export function GetAllTransfers():Promise<main.ErrorReturn>;

export function GetAllTrophies():Promise<main.ErrorReturn>;

export function GetImage(arg1:string):Promise<main.ErrorReturn>;

export function GetNumSaves():Promise<number>;

export function GetNumSeasons():Promise<number>;

export function GetNumSeasonsInSave(arg1:number):Promise<boolean>;

export function GetSaveHomeRankings(arg1:number):Promise<main.ErrorReturn>;

export function GetSavePlayers(arg1:number):Promise<main.ErrorReturn>;

export function GetSavePlayersTotals(arg1:number):Promise<main.ErrorReturn>;

export function GetSaveResults(arg1:number):Promise<main.ErrorReturn>;

export function GetSaveStory(arg1:number):Promise<backend.Story>;

export function GetSaveTransfers(arg1:number):Promise<main.ErrorReturn>;

export function GetSinglePlayer(arg1:number):Promise<main.ErrorReturn>;

export function GetTrophies(arg1:number):Promise<main.ErrorReturn>;

export function Log(arg1:string):Promise<void>;

export function RetrieveSaves():Promise<main.ErrorReturn>;

export function SelectFileParse(arg1:string):Promise<string>;

export function SelectNewTrophyImage(arg1:number):Promise<main.ErrorReturn>;

export function SingleImage(arg1:number):Promise<string>;

export function SingleSave(arg1:number):Promise<main.ErrorReturn>;

export function UpdateSaveStory(arg1:backend.Story):Promise<string>;

export function UploadSaveImage(arg1:number):Promise<main.ErrorReturn>;
