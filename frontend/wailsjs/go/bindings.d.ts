export interface go {
  "main": {
    "KoboService": {
		BuildContentIndex(arg1:Array<Content>):Promise<Content>
		BuildReadwisePayload():Promise<Array<Highlight>|Error>
		CheckReadwiseConfig():Promise<boolean>
		CountDeviceBookmarks():Promise<number>
		DetectKobos():Promise<Array<Kobo>>
		GetReadwiseToken():Promise<string>
		GetSelectedKobo():Promise<Kobo>
		ListDeviceBookmarks():Promise<Array<Bookmark>|Error>
		ListDeviceContent():Promise<Array<Content>|Error>
		NormaliseText(arg1:string):Promise<string>
		OpenDBConnection(arg1:string):Promise<Error>
		PromptForLocalDBPath():Promise<Error>
		SelectKobo(arg1:string):Promise<boolean>
		SendBookmarksToReadwise():Promise<number|Error>
		SetReadwiseToken(arg1:string):Promise<Error>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
