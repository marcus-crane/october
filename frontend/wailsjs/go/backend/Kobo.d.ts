// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {backend} from '../models';
import {map[string]backend} from '../models';

export function CountDeviceBookmarks():Promise<number>;

export function CountDeviceBooks():Promise<number>;

export function FindBookOnDevice(arg1:string):Promise<backend.Content|Error>;

export function ListBookmarksByID(arg1:string):Promise<Array<backend.Bookmark>|Error>;

export function ListBooksOnDevice():Promise<Array<backend.Content>|Error>;

export function ListDeviceBookmarks():Promise<Array<backend.Bookmark>|Error>;

export function ListDeviceContent():Promise<Array<backend.Content>|Error>;

export function BuildContentIndex(arg1:Array<backend.Content>):Promise<map[string]backend.Content>;
