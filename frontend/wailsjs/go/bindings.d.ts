interface go {
  "main": {
    "KoboService": {
		BuildContentIndex(arg1:Array<Content>):Promise<Content>
		BuildReadwisePayload():Promise<Array<Highlight>|Error>
		CountDeviceBookmarks():Promise<number>
		DetectKobos():Promise<Array<Kobo>>
		GetSelectedKobo():Promise<Kobo>
		ListDeviceBookmarks():Promise<Array<Bookmark>|Error>
		ListDeviceContent():Promise<Array<Content>|Error>
		NormaliseText(arg1:string):Promise<string>
		OpenDBConnection(arg1:string):Promise<Error>
		PromptForLocalDBPath():Promise<Error>
		SelectKobo(arg1:string):Promise<boolean>
		SendBookmarksToReadwise():Promise<number|Error>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
