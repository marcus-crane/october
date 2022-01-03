interface go {
  "main": {
    "KoboService": {
		CountDeviceBookmarks():Promise<number>
		DetectKobos():Promise<Array<Kobo>>
		GetSelectedKobo():Promise<Kobo>
		ListDeviceBookmarks():Promise<Array<Bookmark>>
		ListDeviceContent():Promise<Error>
		OpenDBConnection(arg1:string):Promise<Error>
		PromptForLocalDBPath():Promise<Error>
		SelectKobo(arg1:string):Promise<boolean>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
