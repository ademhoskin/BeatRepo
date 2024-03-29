package repo.beat.middleware.model;

import java.util.Date;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;

@Entity
public class BeatZip {
    @Id
    @Column(nullable = false)
    private Long id;

    @Column(nullable = false)
    private String s3Link;

    @Column(nullable = false)
    private String zipName;

    @Column(nullable = false)
    private Long zipSize;

    @Column(nullable = false)
    private Date dateAdded;

    public BeatZip() {
    }

    public BeatZip(String zipName, Long zipSize, String s3Link, Date dateAdded) {
        this.zipName = zipName;
        this.zipSize = zipSize;
        this.s3Link = s3Link;
        this.dateAdded = dateAdded;
    }

    // get set
    public Long getId() {
        return this.id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getZipName() {
        return this.zipName;
    }

    public void setZipName(String zipName) {
        this.zipName = zipName;
    }

    public Long getZipSize() {
        return this.zipSize;
    }

    public void setZipSize(Long zipSize) {
        this.zipSize = zipSize;
    }

    public String getS3Link() {
        return this.s3Link;
    }

    public void setS3Link(String s3Link) {
        this.s3Link = s3Link;
    }

    public Date getDateAdded() {
        return this.dateAdded;
    }

    public void setDateAdded(Date dateAdded) {
        this.dateAdded = dateAdded;
    }

    @Override
    public String toString() {
        return "{" + " id='" + "'" + ", zipName='" + getZipName() + "'" + ", zipSize='" + getZipSize() + "'"
                + ", s3Link='" + getS3Link() + "'" + ", dateAdded='" + getDateAdded() + "'" + "}";
    }
}
