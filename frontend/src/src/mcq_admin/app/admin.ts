export class Admin {
  title: string;
  tabs: Admin.Inner[];
}

export namespace Admin {
  export class Inner {
      name: string;
      url: string;
  };
}
